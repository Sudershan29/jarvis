import React from 'react';
import { createStackNavigator } from '@react-navigation/stack';
import SkillScreen from "../screen/dashboard/Skill"
import SkillCreateScreen from "../screen/skill/CreateSkill"
import SkillShowScreen from "../screen/skill/ShowSkill"

const SkillStack = createStackNavigator();

export const SkillStackScreen = () => (
    <SkillStack.Navigator>
        <SkillStack.Screen name="Skills" component={SkillScreen} />
        <SkillStack.Screen name="SkillCreate" component={SkillCreateScreen} />
        <SkillStack.Screen name="SkillShow" component={SkillShowScreen} />
    </SkillStack.Navigator>
);
