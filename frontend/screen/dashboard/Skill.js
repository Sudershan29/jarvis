import React, { useState, useContext } from "react";
import { ScrollView, View, Text, StyleSheet, TouchableOpacity } from 'react-native';
import { getSkills } from "../../api/Skill";
import Skill from "../../component/Skill";
import { showMessage } from "react-native-flash-message";
import { AuthContext } from "../../context/AuthContext";
import * as Animatable from 'react-native-animatable';
import { useNavigation } from '@react-navigation/native';


export default function SkillScreen() {
    const { userToken } = useContext(AuthContext);
    const [skillsData, setSkillsData] = useState([]);
    const [isButtonClicked, setIsButtonClicked] = useState(false);
    const navigation = useNavigation(); // Access navigation using useNavigation hook


    React.useEffect(() => {
        getSkills(userToken)
            .then(res => {
                if (res.success){
                    if(res.skills?.length > 0)
                        setSkillsData(res.skills)
                }else {
                    showMessage({
                        message: res.message,
                        type: "error",
                    });
                }
            })
            .catch(err => {
                showMessage({
                    message: err.message,
                    type: "error",
                });
            });
    }, []);

    const handleButtonClick = () => {
        setTimeout(() => {
            setIsButtonClicked(isButtonClicked);
        }, 1000);

        setIsButtonClicked(!isButtonClicked);

        navigation.navigate('SkillCreate');
    };

    return (
        <View style={styles.container}>
            <ScrollView>
                {skillsData.map((skill, index) => (
                    <Skill name={skill.name} duration={skill.duration} timepreference={skill.timepreference} categories={skill.categories} key={index} />
                ))}
            </ScrollView>
            <TouchableOpacity style={styles.buttonContainer} onPress={handleButtonClick}>
                <Animatable.View animation={isButtonClicked ? 'fadeOut' : 'fadeIn'} duration={100} style={styles.button}>
                    <Text style={styles.buttonText}>{isButtonClicked ? 'x' : '+'}</Text>
                </Animatable.View>
            </TouchableOpacity>
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        padding: 20,
        backgroundColor: 'white',
    },
    buttonContainer: {
        position: 'absolute',
        bottom: 20,
        right: 20,
        width: 50,
        height: 50,
        borderRadius: 25,
        backgroundColor: '#2196f3',
        alignItems: 'center',
        justifyContent: 'center',
        elevation: 3,
    },
    button: {
        width: '100%',
        height: '100%',
        alignItems: 'center',
        justifyContent: 'center',
    },
    buttonText: {
        color: 'white',
        fontSize: 30,
        fontWeight: 'bold',
    }
});
