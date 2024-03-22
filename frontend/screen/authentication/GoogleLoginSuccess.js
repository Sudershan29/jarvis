import React, { useEffect, useContext } from "react";
import { View, Text, StyleSheet } from 'react-native';
import { useRoute } from '@react-navigation/native';
import { AuthContext } from "../../context/AuthContext";

export default function GoogleLoginSuccess({ navigation }) {
    const route = useRoute();

    const { loginGoogleSuccess } = useContext(AuthContext)

    useEffect(() => {
        if(route.params?.token) {
            loginGoogleSuccess(route.params?.token)
        }
    }, [route.params?.token]);

    return (
        <View style={styles.container}>
            <Text> Authentication Successful </Text>
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
