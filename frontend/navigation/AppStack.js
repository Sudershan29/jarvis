
import React, { useContext } from 'react'

import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { NavigationContainer } from '@react-navigation/native';

import Ionicons from 'react-native-vector-icons/Ionicons';
import HomeScreen from "../screen/dashboard/Home"
import ProfileScreen  from "../screen/dashboard/Profile"
import CalendarScreen from "../screen/dashboard/Calendar"
import CommunityScreen from "../screen/dashboard/Community"
import { SkillStackScreen } from './SkillStack';
import { TaskStackScreen } from './TaskStack';

const Tab = createBottomTabNavigator();

export const AppStack = () => {
    return (
            <Tab.Navigator screenOptions={({ route }) => ({
                tabBarIcon: ({ focused, color, size }) => {
                    let iconName;

                    if (route.name === 'Home')
                        iconName = focused ? 'home' : 'home-outline';
                    else if (route.name === 'Profile')
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
                <Tab.Screen name="Home" options={{ title: 'Welcome' }}>
                    {(props) => <HomeScreen {...props} />}
                </Tab.Screen>
                <Tab.Screen name="SkillMain" component={SkillStackScreen} options={{ headerShown: false }} />
                <Tab.Screen name="Calendar" component={CalendarScreen} />
                <Tab.Screen name="TaskMain" component={TaskStackScreen} options={{ headerShown: false }} />
                <Tab.Screen name="Profile" component={ProfileScreen} />
            </Tab.Navigator>
    )
}