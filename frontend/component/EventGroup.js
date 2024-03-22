import React from "react";
import { View, Text, StyleSheet } from 'react-native';

const EventGroup = ({ group, index }) => {
    const isToday = (date) => {
        const today = new Date();
        return date.getDate() === today.getDate() &&
            date.getMonth() === today.getMonth() &&
            date.getFullYear() === today.getFullYear();
    };

    const getColor = (event) => {
        let color = '#00000080'; // Default color
        if (event.isGoogleCalendarEvent) color = '#2196F3'; // Blue
        if (event.isCancelled) color = '#F44336'; // Red

        // Lighten color if event is in the past
        // const currentTime = new Date();
        // if (new Date(event.startTime) < currentTime) {
        //     color += '80'; // Adding 50% opacity
        // }

        return color;
    };

    const startTime = new Date(group[0].startTime);
    const endTime = new Date(group[group.length - 1].endTime);
    const startTimeStr = `${startTime.getHours()}:${startTime.getMinutes() == 0 ? '00' : startTime.getMinutes()}`;
    const endTimeStr = `${endTime.getHours()}:${endTime.getMinutes() == 0 ? '00' : endTime.getMinutes() }`;

    return (
        <View style={styles.eventGroupContainer} index={index}>
            <Text style={styles.eventTime}>{`${startTimeStr} - ${endTimeStr}`}</Text>
            {group.map((event, index) => (
                <View key={index} style={[styles.event, { backgroundColor: getColor(event) }]}>
                    <Text style={styles.eventText}>{event.name}</Text>
                </View>
            ))}
        </View>
    );
};

const styles = StyleSheet.create({
    eventGroupContainer: {
        // flexDirection: 'row',
        padding: 10,
        flexWrap: 'wrap',
    },
    event: {
        padding: 10,
        borderRadius: 5,
        margin: 2,
    },
    eventText: {
        color: '#fff',
    },
    eventTime: {
        fontSize: 14,
        fontWeight: 'bold',
        marginBottom: 5,
    },
});

export default EventGroup;