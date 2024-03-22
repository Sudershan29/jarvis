import React from "react";
import { View, Text, StyleSheet } from 'react-native';
import { Card } from 'react-native-elements'
import { useNavigation } from '@react-navigation/native';

export default function Task({ name, deadline, scheduled, timepreference, description, duration }) {

    const navigation = useNavigation(); // Access navigation using useNavigation hook

    const cardStyle = () => {
        const now = new Date();
        const deadlineDate = new Date(deadline);
        if (deadline && deadlineDate < now) {
            return styles.pastDeadline;
        } else if (scheduled) {
            return styles.scheduled;
        } else {
            return styles.default;
        }
    };

    function handleButtonClick() {
        navigation.navigate('TaskShow', {
            name: name,
            deadline: deadline, 
            duration: duration,
            scheduled: scheduled,
            timePreferences: timepreference,
        });
    }

    return (
        <Card containerStyle={cardStyle()} onClick={() => { handleButtonClick()}}>
            <Card.Title style={styles.title}>{name}</Card.Title>
            <Card.Divider/>
            { deadline && <Text style={styles.text}> {new Date(deadline).toLocaleString()} </Text>}
        </Card>
    )
}

const styles = StyleSheet.create({
    pastDeadline: {
        backgroundColor: '#FADBD8', // Pale red
        borderRadius: 10,
    },
    scheduled: {
        backgroundColor: '#D5F5E3', // Pale green
        borderRadius: 10,
    },
    default: {
        backgroundColor: '#D6EAF8', // Pale blue
        borderRadius: 10,
    },
    title: {
        color: '#34495E', // Dark blue
        fontSize: 18,
    },
    text: {
        color: '#34495E', // Dark blue
        marginBottom: 10,
    },
});
