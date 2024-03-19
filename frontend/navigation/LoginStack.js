import React, { useEffect, useState } from 'react'
import { createNativeStackNavigator } from '@react-navigation/native-stack';
// import { useLocalSearchParams} from 'expo-router';

import LoginScreen from '../screen/authentication/Login';
import LoginResetScreen from '../screen/authentication/LoginReset';
import GoogleLoginSuccess from '../screen/authentication/GoogleLoginSuccess';

const Stack = createNativeStackNavigator();

export const LoginStack = () => {
    return (
        <Stack.Navigator>
            <Stack.Screen name="Login" component={LoginScreen} />
            <Stack.Screen name="LoginReset" component={LoginResetScreen} />
            <Stack.Screen name="GoogleLoginSuccess" component={GoogleLoginSuccess} />
        </Stack.Navigator>
    )
}