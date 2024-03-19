import React from 'react';
import { View, Text, StyleSheet } from 'react-native';
import ProgressBarMini from './ProgressBarMini';

export default function ProgressBar({ title, progress, subProgresses = [] }) {
    return (
        <View style={styles.container}>
            <Text style={styles.title}>{title}</Text>
            <View style={styles.progressBar}>
                <View style={[styles.progress, { width: `${progress}%` }]}>
                    <Text style={styles.progressText}>{`${progress}%`}</Text>
                </View>
            </View>
            {subProgresses.map((sub, index) => (
                <ProgressBarMini key={index} name={sub.name} value={sub.value} index={index} />
            ))}
        </View>
    );
};

const styles = StyleSheet.create({
    container: {
        padding: 10,
        // backgroundColor: '#f3f3f3',
        borderRadius: 10,
        overflow: 'hidden',
        width: '100%',
        // shadowColor: '#ffffff',
        // shadowOffset: {
        //     width: 0,
        //     height: 4,
        // },
    },
    title: {
        fontSize: 18,
        fontWeight: 'bold',
        marginBottom: 5,
    },
    progressBar: {
        height: 20,
        backgroundColor: '#ddd',
        borderRadius: 10,
        overflow: 'hidden',
    },
    progress: {
        height: '100%',
        backgroundColor: '#2196f3',
        justifyContent: 'center',
        borderRadius: 10,
    },
    progressText: {
        color: '#fff',
        textAlign: 'right',
        paddingRight: 5,
    },
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
