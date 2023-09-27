
import React, { useContext } from 'react'

import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { NavigationContainer } from '@react-navigation/native';

import Ionicons from 'react-native-vector-icons/Ionicons';
import HomeScreen from "../screen/Home"
import ProfileScreen  from "../screen/Profile"
import SkillScreen from "../screen/Skill"
import CalendarScreen from "../screen/Calendar"
import CommunityScreen from "../screen/Community"

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
                    else if (route.name === 'Skill')
                        iconName = focused ? 'football' : 'football-outline';
                    else if (route.name === 'Community')
                        iconName = focused ? 'people' : 'people-outline';

                    return <Ionicons name={iconName} size={size} color={color} />;
                },
                tabBarActiveTintColor: 'black',
                tabBarInactiveTintColor: 'gray',
            })}>
                <Tab.Screen name="Home" options={{ title: 'Welcome' }}>
                    {(props) => <HomeScreen {...props} />}
                </Tab.Screen>
                <Tab.Screen name="Skill" component={SkillScreen} />
                <Tab.Screen name="Calendar" component={CalendarScreen} />
                <Tab.Screen name="Community" component={CommunityScreen} />
                <Tab.Screen name="Profile" component={ProfileScreen} />
            </Tab.Navigator>
    )
}