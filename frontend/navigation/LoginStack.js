
import React from 'react'

import { createNativeStackNavigator } from '@react-navigation/native-stack';

import LoginScreen from '../screen/Login';
import LoginResetScreen from '../screen/LoginReset';

const Stack = createNativeStackNavigator();

export const LoginStack = () => {
    return (
        <Stack.Navigator initialRouteName="Login">
            <Stack.Screen name="Login" component={LoginScreen} />
            <Stack.Screen name="LoginReset" component={LoginResetScreen} />
        </Stack.Navigator>
    )
}