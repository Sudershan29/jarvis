import React, { useContext } from "react";
import { View, Text, StyleSheet } from 'react-native';
import TimePreference from "../../component/TimePreference";
import { showMessage } from "react-native-flash-message";
import { AuthContext } from "../../context/AuthContext";

export default function SkillShowScreen({ route, navigation }) {
    const { userToken } = useContext(AuthContext);

    const { skillId, name, duration, timePreferences } = route.params;

    return (
        <View style={styles.container}>
            <Text>Name: Sample { name }</Text>
            <Text>Duration: { duration } </Text>
            <TimePreference disableClick={true} timePreferences={timePreferences} setTimePreferences={()=> {}} />
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
