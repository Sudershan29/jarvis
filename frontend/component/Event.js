import React from "react";
import { View, Text, StyleSheet } from 'react-native';
import EventGroup from "./EventGroup";

const Event = ({ events, heading }) => {
    const groupEvents = () => {
        const grouped = [];
        let lastEvent = null;

        events.forEach(event => {
            const startTime = new Date(event.startTime);
            const endTime = lastEvent ? new Date(lastEvent.endTime) : null;

            if (lastEvent && (startTime - endTime) <= 900000) { // 15 minutes in milliseconds
                grouped[grouped.length - 1].push(event);
            } else {
                grouped.push([event]);
            }

            lastEvent = event;
        });

        return grouped;
    };

    const renderEventGroups = () => {
        const eventGroups = groupEvents();

        return eventGroups.map((group, index) => {
            return (
                <EventGroup group={group} index={index}/>
            );
        });
    };

    return (
        <View style={styles.container}>
            <Text style={{fontWeight: 'bold', fontSize: 20}} >{heading}</Text>
            {renderEventGroups()}
        </View>
    );
};

const styles = StyleSheet.create({
    container: {
        flex: 1,
        padding: 10,
    },
});

export default Event;

