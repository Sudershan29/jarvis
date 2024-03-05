import React from 'react';
import { createStackNavigator } from '@react-navigation/stack';
import TaskScreen from "../screen/dashboard/Task"
import TaskCreateScreen from "../screen/task/CreateTask"
import TaskShowScreen from "../screen/task/ShowTask"

const TaskStack = createStackNavigator();

export const TaskStackScreen = () => (
    <TaskStack.Navigator>
        <TaskStack.Screen name="Tasks" component={TaskScreen} />
        <TaskStack.Screen name="TaskCreate" component={TaskCreateScreen} />
        <TaskStack.Screen name="TaskShow" component={TaskShowScreen} />
    </TaskStack.Navigator>
);
