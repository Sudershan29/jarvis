
import React, { useContext } from 'react'

import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { NavigationContainer } from '@react-navigation/native';

import Ionicons from 'react-native-vector-icons/Ionicons';
import HomeScreen from "../screen/dashboard/Home"
import CalendarScreen from "../screen/dashboard/Calendar"
import CommunityScreen from "../screen/dashboard/Community"
import { SkillStackScreen } from './SkillStack';
import { TaskStackScreen } from './TaskStack';
import { ProfileStackScreen } from './ProfileStack';
import GoogleCalendarConnect from '../screen/authentication/GoogleCalendarConnect';

const Tab = createBottomTabNavigator();

export const AppStack = () => {
    return (
            <Tab.Navigator screenOptions={({ route }) => ({
                tabBarIcon: ({ focused, color, size }) => {
                    let iconName;

                    if (route.name === 'Home')
                        iconName = focused ? 'home' : 'home-outline';
                    else if (route.name === 'ProfileMain')
                        iconName = focused ? 'person-circle' : 'person-circle-outline';
                    else if (route.name === 'Calendar')
                        iconName = focused ? 'calendar' : 'calendar-outline';
                    else if (route.name === 'SkillMain')
                        iconName = focused ? 'football' : 'football-outline';
                    else if (route.name === 'TaskMain')
                        iconName = focused ? 'clipboard' : 'clipboard-outline';

                    return <Ionicons name={iconName} size={size} color={color} />;
                },
                tabBarActiveTintColor: 'black',
                tabBarInactiveTintColor: 'gray',
            })}>
                <Tab.Screen name="Home" component={HomeScreen} options={{ title: 'Welcome' }} />
                <Tab.Screen name="TaskMain" component={TaskStackScreen} options={{ headerShown: false, title: 'Tasks' }} />
                <Tab.Screen name="Calendar" component={CalendarScreen} />
                <Tab.Screen name="SkillMain" component={SkillStackScreen} options={{ headerShown: false, title: 'Skills' }} />
                <Tab.Screen name="ProfileMain" component={ProfileStackScreen} options={{ headerShown: false, title: 'Profile' }} />
                <Tab.Screen name="GoogleCalendarConnect" component={GoogleCalendarConnect} options={{ tabBarButton: () => null }} />
            </Tab.Navigator>
    )
}