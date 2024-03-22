import React from 'react';
import { createStackNavigator } from '@react-navigation/stack';
import ProfileScreen from '../screen/dashboard/Profile';
import GoogleCalendarConnect from '../screen/authentication/GoogleCalendarConnect';

const ProfileStack = createStackNavigator();

export const ProfileStackScreen = () => (
    <ProfileStack.Navigator>
        <ProfileStack.Screen name="Profile" component={ProfileScreen} />
        {/* <ProfileStack.Screen name="GoogleCalendarConnect" component={GoogleCalendarConnect} /> */}
    </ProfileStack.Navigator>
);
