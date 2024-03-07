import React from 'react';
import { View, Text, StyleSheet } from 'react-native';

export default function ProgressBarMini({ name, value, index }) {
    return (
        <View key={index} style={styles.subProgressBar}>
            <Text style={styles.subTitle}>{name}</Text>
            <View style={[styles.subProgress, { width: `${value}%` }]}>
                <Text style={styles.subProgressText}>{`${value}%`}</Text>
            </View>
        </View>
    )
}

const styles = StyleSheet.create({
    subProgressBar: {
        marginTop: 10,
    },
    subTitle: {
        fontSize: 14,
        marginBottom: 3,
    },
    subProgress: {
        height: 10,
        backgroundColor: '#bbb',
        borderRadius: 5,
        overflow: 'hidden',
    },
    subProgressText: {
        color: '#fff',
        textAlign: 'right',
        paddingRight: 2,
        fontSize: 10,
    },
});