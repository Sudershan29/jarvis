
import React from "react";
import { Button, View, Text, StyleSheet } from 'react-native';

export default function CommunityScreen () {
    return (
        <View style={styles.container}>
            <Text>Community Screen</Text>
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#fff',
        alignItems: 'center',
        justifyContent: 'center',
    },
});
