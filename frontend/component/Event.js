import React from "react";
import { View, Text, StyleSheet } from 'react-native';
import EventGroup from "./EventGroup";

const Event = ({ events, heading, isDate }) => {
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
                <EventGroup key={index} group={group} index={index}/>
            );
        });
    };

    return (
        <View style={styles.container}>
                {!isDate && <Text style={{fontWeight: 'bold', fontSize: 20}} >{heading}</Text>}
                {isDate && <Text style={{fontWeight: 'bold', fontSize: 20}} >{heading?.day}, {heading?.date + " " + heading?.month}</Text>}
                {events.length === 0 ?
                    <View style={{ alignItems: 'center' }}>
                        <Text style={{padding: 20}}>No upcoming events for today</Text> 
                    </View> 
                    : renderEventGroups()}
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

