
import React, { useContext } from 'react'
import { AuthContext } from "../context/AuthContext"
import { ActivityIndicator } from 'react-native';
import { NavigationContainer } from '@react-navigation/native';

import { AppStack } from './AppStack';
import { LoginStack } from "./LoginStack"

export const AppNavigation = () => {
    const { isLoggedIn, isLoading } = useContext(AuthContext)

    if(isLoading)
        return (<ActivityIndicator/>)

    return(
        <NavigationContainer>
         { isLoggedIn() ? <AppStack /> : <LoginStack/> }
        </NavigationContainer>
    )
}