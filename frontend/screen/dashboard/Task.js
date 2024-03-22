
import React, { useState, useContext } from "react";
import { ScrollView, View, Text, StyleSheet, TouchableOpacity } from 'react-native';
import { getTasks } from "../../api/Task";
import Task from "../../component/Task"
import { showMessage } from "react-native-flash-message";
import { AuthContext } from "../../context/AuthContext"
import * as Animatable from 'react-native-animatable';
import { useNavigation } from '@react-navigation/native';

export default function TaskScreen() {
    const { userToken } = useContext(AuthContext)
    const [TasksData, setTasksData] = useState([])
    const [isButtonClicked, setIsButtonClicked] = useState(false);
    const navigation = useNavigation();

    React.useEffect(() => {
        getTasks(userToken)
            .then(res => { 
                if(res.success)
                    if (res.tasks?.length > 0){
                        setTasksData(res.tasks)
                    }
                else{
                    showMessage({
                        message: res.message,
                        type: "error",
                    })
                }
             })
            .catch(err => {
                showMessage({
                    message: err.message,
                    type: "error",
                })
            })
    }, [])

    const handleButtonClick = () => {
        setTimeout(() => {
            setIsButtonClicked(isButtonClicked);
        }, 1000);

        setIsButtonClicked(!isButtonClicked);

        navigation.navigate('TaskCreate');
    };

    return (
        <View style={styles.container}>
            <ScrollView>
                {TasksData.map((task, index) => (
                    <Task 
                        id={task.id}
                        name={task.name} 
                        deadline={task.deadline} 
                        key={index}
                        description={task.description}
                        duration={task.duration}
                        categories={task.categories}
                        timepreference={task.timepreference} />
                ))}
            </ScrollView>
            <TouchableOpacity style={styles.buttonContainer} onPress={handleButtonClick}>
                <Animatable.View animation={isButtonClicked ? 'fadeOut' : 'fadeIn'} duration={300} style={styles.button}>
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
        backgroundColor: '#2196f3', // Deep blue color
        alignItems: 'center',
        justifyContent: 'center',
        elevation: 3, // Add some elevation for Android shadow
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
