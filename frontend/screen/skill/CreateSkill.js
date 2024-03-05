import React, { useState, useContext } from "react";
import { Button, View, Text, TextInput, StyleSheet } from 'react-native';
import TimePreference from "../../component/TimePreference";
import { showMessage } from "react-native-flash-message";
import { useNavigation } from '@react-navigation/native';
import { AuthContext } from "../../context/AuthContext";
import { createSkill } from "../../api/Skill";

export default function SkillCreateScreen() {
    const { userToken } = useContext(AuthContext);
    const [name, setName] = useState('');
    const [level, setLevel] = useState('');
    const [duration, setDuration] = useState(0);
    const [showDetails, setShowDetails] = useState(false);
    const [categories, setCategories] = useState('');
    const [timePreferences, setTimePreferences] = useState([]);
    const navigation = useNavigation();

    const handleSubmit = () => {
        const skillData = {
            name: name,
            duration: duration,
            level: showDetails ? level : '',
            categories: showDetails ? [categories] : [],
            timepreferences: showDetails ? timePreferences : [],
        };
        createSkill(userToken, skillData)
            .then(res => {
                if (res.success) {
                    showMessage({
                        message: "Skill created successfully",
                        type: "success",
                    });
                    navigation.navigate('Skills');
                } else {
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
                    <TimePreference timePreferences={timePreferences} setTimePreferences={setTimePreferences} />
                    <TextInput
                        style={styles.input}
                        placeholder="Categories"
                        value={categories}
                        onChangeText={text => setCategories(text)}/>

                    <TextInput
                        style={styles.input}
                        placeholder="Level"
                        value={level}
                        onChangeText={text => setLevel(text)}
                    />
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
