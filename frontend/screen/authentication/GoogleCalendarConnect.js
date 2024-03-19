import React, { useEffect, useContext } from "react";
import { View, Text, StyleSheet } from 'react-native';
// import { useRoute } from '@react-navigation/native';
// import { AuthContext } from "../../context/AuthContext";

export default function GoogleCalendarConnect({ navigation }) {
    // const route = useRoute();

    // const { calendarConnectSuccess } = useContext(AuthContext)
    // console.log("GoogleCalendarConnect", route.params?.success)

    // useEffect(() => {
    //     if (route.params?.success) {
    //         console.log("Redirecting to ProfileMain", route.params?.success)
    //         calendarConnectSuccess(() => { navigation.navigate('Profile') })
    //     }
    // }, []);

    return (
        <View style={styles.container}>
            <Text> Google Calendar Successful </Text>
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
