import React from "react";
import { View, Text, StyleSheet, TouchableOpacity } from 'react-native';

const ProposalGroup = ({ proposals, cancel }) => {
    const formatDate = (dateString) => {
        const options = { year: 'numeric', month: 'long', day: 'numeric', hour: '2-digit', minute: '2-digit' };
        return new Date(dateString).toLocaleDateString(undefined, options);
    };

    const getStatusPillColor = (status) => {
        switch (status) {
            case 'pending':
                return '#2196F3'; // Blue
            case 'cancelled':
                return '#F44336'; // Red
            case 'done':
                return '#4CAF50'; // Green
            default:
                return '#757575'; // Default grey
        }
    };

    return (
        <View style={styles.container}>
            <Text style={{fontSize: 20, fontWeight: 'bold', marginBottom: 10}}>Calendar Events for this week</Text>
            {proposals.map((proposal, index) => (
                <View key={index} style={styles.row}>
                    {index > 0 && <View style={{ height: 1, backgroundColor: '#e0e0e0', marginVertical: 10 }}></View>}
                    <Text style={styles.cell}>{proposal.id}</Text>
                    <Text style={styles.cell}>{formatDate(proposal.scheduledFor)}</Text>
                    <View style={[styles.statusPill, { backgroundColor: getStatusPillColor(proposal.status) }]}>
                        <Text style={styles.statusText}>{proposal.status}</Text>
                    </View>
                    {proposal.status === 'pending' && (
                        <TouchableOpacity onPress={() => cancel(proposal.id)} style={styles.cancelButton}>
                            <Text style={styles.cancelButtonText}>Cancel</Text>
                        </TouchableOpacity>
                    )}
                </View>
            ))}
        </View>
    );
};

const styles = StyleSheet.create({
    container: {
        padding: 10,
    },
    row: {
        flexDirection: 'row',
        alignItems: 'center',
        marginBottom: 10,
    },
    cell: {
        marginRight: 10,
    },
    statusPill: {
        borderRadius: 15,
        paddingVertical: 5,
        paddingHorizontal: 10,
        marginRight: 10,
    },
    statusText: {
        color: '#fff',
    },
    cancelButton: {
        backgroundColor: '#F44336',
        borderRadius: 5,
        padding: 5,
    },
    cancelButtonText: {
        color: '#fff',
    },
});

export default ProposalGroup;
