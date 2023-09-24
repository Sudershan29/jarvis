
import React from "react";
import { Button, View, Text, StyleSheet } from 'react-native';

export default function LoginResetScreen({ navigation }) {
    return (
        <View style={styles.container}>
            <Text>Forgot Password Screen </Text>
            <Button
                title="Go to Login"
                onPress={() => navigation.navigate('Login')}
            />
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#fff',
        alignItems: 'center',
        justifyContent: 'center',
    },
});
