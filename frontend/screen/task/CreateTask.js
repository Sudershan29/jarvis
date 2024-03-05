import React, { useState, useContext } from "react";
import { Button, View, Text, TextInput, StyleSheet } from 'react-native';
import TimePreference from "../../component/TimePreference";
import { showMessage } from "react-native-flash-message";
import { useNavigation } from '@react-navigation/native';
import { AuthContext } from "../../context/AuthContext";
import DateTimePicker from '@react-native-community/datetimepicker';

export default function TaskCreateScreen() {
    const { userToken } = useContext(AuthContext);
    const [name, setName] = useState('');
    const [description, setDescription] = useState('');
    const [deadline, setDeadline] = useState('');
    const [duration, setDuration] = useState(0);
    const [showDetails, setShowDetails] = useState(false);
    const [categories, setCategories] = useState('');
    const [timePreferences, setTimePreferences] = useState([]);
    const navigation = useNavigation();

    const handleSubmit = () => {
        const taskData = {
            name: name,
            description: description,
            duration: duration,
            deadline: deadline,
            categories: showDetails && categories.length ? [categories] : [],
            timepreferences: showDetails ? timePreferences : [],
        };
        console.log(taskData);
        // createSkill(userToken, skillData)
        //     .then(res => {
        //         if (res.success) {
        //             showMessage({
        //                 message: "Skill created successfully",
        //                 type: "success",
        //             });
        //             navigation.navigate('SkillMain');
        //         } else {
        //             showMessage({
        //                 message: res.message,
        //                 type: "error",
        //             });
        //         }
        //     })
        //     .catch(err => {
        //         showMessage({
        //             message: err.message,
        //             type: "error",
        //         });
        //     });
    };

    const handleDeadlineChange = (event, selectedDate) => {
        const currentDate = selectedDate || new Date();
        setDeadline(currentDate);
    };

    return (
        <View style={styles.container}>
            <TextInput
                style={styles.input}
                placeholder="Name"
                value={name}
                onChangeText={text => setName(text)}
            />

            <TextInput
                style={styles.input}
                placeholder="Description"
                value={description}
                onChangeText={text => setDescription(text)}
            />
            
            <TextInput
                style={styles.input}
                placeholder="Duration"
                value={duration.toString()}
                keyboardType="numeric"
                onChangeText={text => {
                    const parsedValue = parseInt(text);
                    if (!isNaN(parsedValue)) {
                        setDuration(parsedValue);
                    } else {
                        setDuration(0);
                    }
                }}
            />

            <Button title={showDetails ? "Hide Optional Preferences" : "Add Optional Preferences"} onPress={() => setShowDetails(!showDetails)} />
            {showDetails && (
                <View>
                    <DateTimePicker
                        style={styles.input}
                        placeholder="Deadline"
                        value={deadline}
                        mode="datetime"
                        minimumDate={new Date().toISOString().split('T')[0]}
                        onChange={handleDeadlineChange}
                    />

                    <TimePreference timePreferences={timePreferences} setTimePreferences={setTimePreferences} />
                    <TextInput
                        style={styles.input}
                        placeholder="Categories"
                        value={categories}
                        onChangeText={text => setCategories(text)} />

                </View>
            )}
            <Button title="Submit" onPress={handleSubmit} />
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#fff',
        alignItems: 'center',
        justifyContent: 'center',
        padding: 20,
    },
    input: {
        height: 40,
        borderColor: 'gray',
        borderWidth: 1,
        margin: 10,
        padding: 10,
        width: '80%',
    },
});
