import React, { useState, useEffect, useContext } from "react";
import { Button, View, Text, StyleSheet, ScrollView } from 'react-native';
import DatePill from "../../component/DatePill";
import Event from '../../component/Event';
import { AuthContext } from "../../context/AuthContext";
import { getEvents } from "../../api/Calendar";

export default function CalendarScreen({ startOfDay }) {
    const today = new Date();
    today.setHours(0, 0, 0, 0);

    const { userToken } = useContext(AuthContext);
    const [currentDateRange, setCurrentDateRange] = useState([]);
    const [activeDateIndex, setActiveDateIndex] = useState(0);
    const [currentDate, setCurrentDate] = useState(today);
    const [currEvents, setCurrEvents] = useState([]);

    useEffect(() => {
        generateDateRange(0);
    }, []);

    useEffect(() => {
        const fetchData = async () => {
            let dayAfterCurrentDate = new Date(currentDate);
            dayAfterCurrentDate.setDate(currentDate.getDate() + 1); // Add 24 hours to the current date
            const events = await getEvents(userToken, currentDate.toISOString(), dayAfterCurrentDate.toISOString());
            setCurrEvents(events);
        };

        fetchData();
    }, [activeDateIndex]);

    const generateDateRange = (offset) => {
        let dates = [];
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
            <View style={styles.dateScrollContainer}>
                <ScrollView horizontal={true} showsHorizontalScrollIndicator={false}>
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
            </View>
            <View style={styles.eventContainer}>
                <ScrollView>
                    <Event 
                        heading={currentDate === new Date().getDate() ? "Today" : currentDateRange[activeDateIndex]}
                        isDate={currentDate !== new Date().getDate()}
                        events={currEvents}
                    />
                </ScrollView>
            </View>
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        flexGrow: 1,
        backgroundColor: '#fff',
        flexDirection: 'column',
        padding: 10,
    },
    dateScrollContainer: {
        height: 100, // Set a fixed height for the date scroll container
    },
    eventContainer: {
        flex: 1, // Allow this container to expand and fill the space
        // paddingLeft: 10,
        // paddingRight: 10,
    },
});
