
import React, {useContext} from "react";
import { Button, View, Text, StyleSheet } from 'react-native';

import { AuthContext } from "../context/AuthContext"

export default function ProfileScreen () {
    const { logout } = useContext(AuthContext)

    return (
        <View style={styles.container}>
            <Text>Profile Screen</Text>
            <Button
                title="Logout"
                onPress={() => logout()}
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
