
import React, { useContext } from 'react'
import { View } from 'react-native';
import { AuthContext } from "../context/AuthContext"
import { ActivityIndicator } from 'react-native';
import { NavigationContainer } from '@react-navigation/native';
import FlashMessage from "react-native-flash-message";

import { AppStack } from './AppStack';
import { LoginStack } from "./LoginStack"

const config = {
    screens: {
        ProfileMain: {
            screens: {
                GoogleCalendarConnect: 'GoogleCalendarConnect',
                Profile: 'Profile',
            },
        },
        GoogleLoginSuccess: 'googleLoginSuccess',
        CalendarDone: 'CalendarDone'
    },
};

const linking = {
    config,
};

export const AppNavigation = () => {
    const { isLoggedIn, isLoading } = useContext(AuthContext)

    if(isLoading)
        return (<ActivityIndicator/>)

    return(
        <NavigationContainer linking={linking} fallback={<Text>Loading...</Text>}>
            <View style={{ flex: 1 }}>
                { isLoggedIn() ? <AppStack /> : <LoginStack/> }
                <FlashMessage position="top" />
            </View>
        </NavigationContainer>
    )
}