import React, { useContext } from "react";
import { View, Text, StyleSheet } from 'react-native';
import { showMessage } from "react-native-flash-message";
import { AuthContext } from "../../context/AuthContext";
import TimePreference from "../../component/TimePreference";

export default function TaskShowScreen({ route, navigation }) {
    const { userToken } = useContext(AuthContext);
    const { taskId, name, deadline, duration, scheduled, timePreferences } = route.params;
    const deadlineDate = new Date(deadline);

    return (
        <View style={styles.container}>
            <Text>Name: Sample {name}</Text>
            <Text>Deadline: {deadlineDate.toLocaleString()} </Text>
            <Text>Total Duration: {duration} hours </Text>
            <TimePreference disableClick={true} timePreferences={timePreferences} setTimePreferences={() => { }} />
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#fff',
        alignItems: 'center',
        justifyContent: 'center',
        padding: 20,
    },
});
