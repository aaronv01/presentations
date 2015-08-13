package main

import (
	"database/sql"
	"fmt"
	"time"
)

//show main OMIT
func main() {
	// Server Authentication stuff...

	for {
		Remind(Retrieve(appUID)) // HL
		time.Sleep(15 * time.Second)
	}
}

//end show main OMIT

//show signature OMIT
func Retrieve(appId string) <-chan *Reminder { // HL
	// ...
}

//end show signature OMIT
func Retrieve(appId string) <-chan *Reminder {
	//show query OMIT
	log("Retrieving reminders from database")
	db := db()
	defer db.Close()

	st, err := db.Prepare("select * from userreminder where reminderstatus=0 and appuid=?")
	if err != nil {
		fmt.Print(err)
		return nil
	}
	//end show query
	rows, err := st.Query(appId)
	if err != nil {
		fmt.Print(err)
		return nil
	}
	//show parser OMIT
	c := make(chan *Reminder) // HL
	go func() {
		for rows.Next() {
			var dayBefore, hourBefore sql.RawBytes
			reminder := new(Reminder)
			if err := rows.Scan(&reminder.Id, &reminder.UID, &reminder.EndUserId,
				&reminder.ScheduleItemId, &reminder.DateStr, &reminder.Message,
				&reminder.Title, &dayBefore, &hourBefore,
				&reminder.CustomBefore, &reminder.ReminderStatus, &reminder.AppUID); err != nil {
				panic(err)
			}

			c <- reminder // HL
		}
		close(c) // HL
	}()
	return c
	//end show parser OMIT
}

//show reminder OMIT
func Remind(c <-chan *Reminder) { // HL
	for reminder := range c { // HL
		go func(reminder *Reminder) { // HL
			for {
				t := time.Now().In(location)
				duration := t.Sub(reminder.Date)
				// Time validations...
				if duration.Hours() >= -1 && reminder.HourBefore {
					log(fmt.Sprintln("1 hour reminder"))
					reminder.HourBefore = false
					reminder.Message = "Don't forget your doctor's appointment for later today!"
					sendAlert(reminder) // HL
				}

				if duration.Minutes() >= -float64(reminder.CustomBefore) && !reminder.DayBefore && !reminder.HourBefore {
					log(fmt.Sprintf("Done with reminder %v", reminder.Id))
					break
				}
				// Estimate next alert onto waitingTime variable..
				time.Sleep(waitingTime)
			}
		}(reminder)
	}
}

//end show reminder OMIT
