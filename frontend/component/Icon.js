import React from 'react';
import { View, Text, StyleSheet, TouchableOpacity, Image } from 'react-native';
import Ionicons from 'react-native-vector-icons/Ionicons';

const Icon = ({ name, execute, image, key }) => {
    return (
        <View style={styles.iconContainer}>
            <TouchableOpacity key={key} onPress={execute} style={styles.container}>
                <Ionicons name={image} style={styles.icon} size={30} />
            </TouchableOpacity>
            <Text style={styles.text}>{name}</Text>
        </View>
    );
};

const styles = StyleSheet.create({
    container: {
        backgroundColor: '#f2f2f2',
        borderRadius: 5,
        alignItems: 'center',
        justifyContent: 'center',
        width: 60,
        height: 60,
    },
    iconContainer: {
        alignItems: 'center',
        paddingRight: 15,
        paddingLeft: 15, 
        paddingTop: 15,
        // justifyContent: 'center',
    },
    icon: {
        marginBottom: 5,
    },
    text: {
        fontSize: 10,
        textAlign: 'center',
        width: 60,
    },
});

export default Icon;
