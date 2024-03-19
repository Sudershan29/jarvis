import React, { useState, useEffect, useContext } from "react";
import { Button, View, Text, StyleSheet, ScrollView } from 'react-native';
import DatePill from "../../component/DatePill";
import Event from '../../component/Event';
import { AuthContext } from "../../context/AuthContext";
import { getEvents } from "../../api/Calendar";

export default function CalendarScreen() {

    const { userToken } = useContext(AuthContext);
    const [currentDateRange, setCurrentDateRange] = useState([]);
    const [activeDateIndex, setActiveDateIndex] = useState(0);
    const [currentDate, setCurrentDate] = useState(new Date());
    const [currEvents, setCurrEvents] = useState([]);

    useEffect(() => {
        generateDateRange(0);
    }, []);

    useEffect(() => {
        const fetchData = async () => {
            const dayAfterCurrentDate = new Date(currentDate);
            dayAfterCurrentDate.setDate(dayAfterCurrentDate.getDate() + 1); // Add 24 hours to the current date
            const events = await getEvents(userToken, currentDate.toISOString(), dayAfterCurrentDate.toISOString());
            setCurrEvents(events);
        };

        fetchData();
    }, [activeDateIndex]);

    // Make an API call based on currentDate

    const generateDateRange = (offset) => {
        let dates = [];
        let today = new Date();
        today.setDate(today.getDate() + offset); // Adjust today based on offset
        for (let i = -3; i <= 3; i++) { // Generate 7 days range with today in the center
            let date = new Date(today);
            date.setDate(date.getDate() + i);
            dates.push({
                dateObj: date,
                date: date.getDate(),
                day: date.toLocaleString('en-us', { weekday: 'short' }),
                month: date.toLocaleString('en-us', { month: 'short' }),
            });
        }
        setCurrentDateRange(dates);
        setActiveDateIndex(3);
    };

    const changeActiveDate = (index) => {
        setCurrentDate(currentDateRange[index].dateObj);
        setActiveDateIndex(index);
    };

    return (
        <View style={styles.container}>
            {/* <Button title="<" onPress={() => generateDateRange(-1)} /> */}
            <ScrollView horizontal={true} style={styles.dateScroll} showsHorizontalScrollIndicator={false}>
                {currentDateRange.map((dateInfo, index) => (
                    <DatePill
                        key={index}
                        date={dateInfo.date}
                        day={dateInfo.day}
                        month={dateInfo.month}
                        highlighted={index === activeDateIndex} // Highlight the active date
                        today={index === 3} // Highlight today
                        onPressFunc={() => {changeActiveDate(index)}} // Set the clicked date as active
                    />
                ))}
            </ScrollView>
            <View style={styles.eventContainer}>
                <Event 
                    heading={currentDate === new Date().getDate() ? "Today" : currentDateRange[activeDateIndex]}
                    isDate={currentDate !== new Date().getDate()}
                    events={currEvents}
                    // events={[{ name: "Breakfast", isCancelled: false, startTime: "2023-03-08T08:00:00", endTime: "2023-03-08T09:00:00"}, 
                    //  { name: "Robert <> Sudershan", isCancelled: false, startTime: "2023-03-08T09:00:00", endTime: "2023-03-08T10:00:00", isGoogleCalendarEvent: true },
                    //  { name: "Fake Event", isCancelled: true, startTime: "2023-03-08T11:30:00", endTime: "2023-03-08T12:30:00" },
                    //  { name: "Badminton", isCancelled: false, startTime: "2023-03-08T18:30:00", endTime: "2023-03-08T20:30:00" }]}
                />
            </View>
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#fff',
        flexDirection: 'column',
        padding: 10,
    },
    dateScroll: {
        flex: 1,
        // marginHorizontal: 50,
    },
    eventContainer: {
        flex: 6,
        // paddingLeft: 10,
        // paddingRight: 10,
    },
});
