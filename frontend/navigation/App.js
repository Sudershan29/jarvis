
import React, { useContext } from 'react'
import { View } from 'react-native';
import { AuthContext } from "../context/AuthContext"
import { ActivityIndicator } from 'react-native';
import { NavigationContainer } from '@react-navigation/native';
import FlashMessage from "react-native-flash-message";

import { AppStack } from './AppStack';
import { LoginStack } from "./LoginStack"

export const AppNavigation = () => {
    const { isLoggedIn, isLoading } = useContext(AuthContext)

    if(isLoading)
        return (<ActivityIndicator/>)

    return(
        <NavigationContainer>
            <View style={{ flex: 1 }}>
                { isLoggedIn() ? <AppStack /> : <LoginStack/> }
                <FlashMessage position="top" />
            </View>
        </NavigationContainer>
    )
}