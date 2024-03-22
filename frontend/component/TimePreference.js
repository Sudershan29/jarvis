import React, { useState } from 'react';
import { View, StyleSheet, TouchableOpacity, Text } from 'react-native';

const daysOfTheWeek = [
    { name: 'Sun', value: 0, fullName: 'sunday'},
    { name: 'Mon', value: 1, fullName: 'monday'},
    { name: 'Tue', value: 2, fullName: 'tuesday'},
    { name: 'Wed', value: 3, fullName: 'wednesday'},
    { name: 'Thu', value: 4, fullName: 'thursday'},
    { name: 'Fri', value: 5, fullName: 'friday'},
    { name: 'Sat', value: 6, fullName: 'saturday'},
]

export default function TimePreference({ timePreferences, setTimePreferences, disableClick}) {

    const handleDayClick = (day) => {
        if (disableClick) return;

        if (timePreferences.includes(day.fullName)) {
            setTimePreferences(timePreferences.filter(d => d !== day.fullName));
        } else {
            setTimePreferences([...timePreferences, day.fullName]);
        }
    };

    return (
        <View style={styles.container}>
            {daysOfTheWeek.map(day => (
                <TouchableOpacity
                    key={day.value}
                    style={[styles.dayCircle, timePreferences.includes(day.fullName) && styles.selectedDay]}
                    onPress={() => handleDayClick(day)}
                >
                    <Text style={styles.buttonText}>{day.name}</Text>
                </TouchableOpacity>
            ))}
        </View>
    );
};

const styles = StyleSheet.create({
    container: {
        flexDirection: 'row',
        justifyContent: 'space-around',
        marginTop: 20,
    },
    dayCircle: {
        width: 50,
        height: 50,
        borderRadius: 35,
        padding: 10,
        backgroundColor: '#E0E0E0',
        alignItems: 'center',
        justifyContent: 'center',
    },
    selectedDay: {
        backgroundColor: '#2196f3',
    },
});

