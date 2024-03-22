import React, { useState, useContext, useEffect } from "react";
import { View, Text, StyleSheet, Button } from 'react-native';
import { showMessage } from "react-native-flash-message";
import { AuthContext } from "../../context/AuthContext";
import TimePreference from "../../component/TimePreference";
import { getProposals, cancelProposal } from "../../api/Task";
import ProposalGroup from "../../component/ProposalGroup";

export default function TaskShowScreen({ route, navigation }) {
    const { userToken } = useContext(AuthContext);
    const { id, name, deadline, duration, scheduled, timePreferences } = route.params;
    const deadlineDate = new Date(deadline);
    const [proposals, setProposals] = useState([]);

    useEffect(() => {
        getProposals(userToken, id).then(res => {
            if (res.success) {
                setProposals(res.proposals);
            } else {
                showMessage({
                    message: res.message,
                    type: "error",
                });
            }
        })
    }, [])


    return (
        <View style={styles.container}>
            <Text>Name: Sample {name}</Text>
            <Text>Deadline: {deadlineDate.toLocaleString()} </Text>
            <Text>Total Duration: {duration} hours </Text>
            <TimePreference disableClick={true} timePreferences={timePreferences} setTimePreferences={() => { }} />
            <ProposalGroup proposals={proposals} cancel={(proposalId) => {cancelProposal(id, proposalId)}} />
            <Button title="Mark as Done" onPress={()=> {}}/>
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
});
