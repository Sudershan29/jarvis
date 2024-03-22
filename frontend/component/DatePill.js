import React from "react";
import { View, Text, StyleSheet } from 'react-native';

export default function DatePill({ date, day, month, highlighted, onPressFunc, today }) {
    return (
        <View style={[styles.container, highlighted ? styles.highlighted : null]} onClick={onPressFunc}>
            <Text style={{ fontSize: 12, textTransform: 'uppercase', fontWeight: highlighted ? 'bold' : 'normal' }}>{day}</Text>
            <View style={styles.divider}></View>
            <Text style={{ fontSize: 10, fontWeight: 'normal', whiteSpace: 'nowrap' }}>{date} {month}</Text>
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#fff',
        alignItems: 'center',
        justifyContent: 'center',
        borderWidth: 1,
        borderColor: '#ccc',
        borderRadius: 20,
        margin: 5,
        padding: 10,
        height: 75,
    },
    highlighted: {
        backgroundColor: '#def1ff',
    },
    divider: {
        width: '100%',
        height: 1,
        backgroundColor: '#ccc',
        marginVertical: 5,
    },
});
