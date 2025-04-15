// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title AuditTrail
 * @dev Smart contract for recording immutable audit logs for AI interactions
 */
contract AuditTrail {
    // Struct to represent an audit log entry
    struct LogEntry {
        address recorder;      // Address that recorded this log
        string userIdentifier; // External user ID
        string actionType;     // Type of action (completion, chat, etc.)
        string contentHash;    // Hash of the content (request/response)
        string metadata;       // Additional metadata (JSON string)
        uint256 timestamp;     // Timestamp when the log was recorded
    }

    // Array to store all log entries
    LogEntry[] private logs;

    // Mapping from user identifier to their log indices
    mapping(string => uint256[]) private userLogs;
    
    // Mapping from content hash to log index
    mapping(string => uint256) private hashToLogIndex;

    // Events
    event LogRecorded(string userIdentifier, string actionType, string contentHash, uint256 timestamp);

    /**
     * @dev Record a new audit log entry
     * @param userIdentifier Identifier for the user
     * @param actionType Type of action (completion, chat, etc.)
     * @param contentHash Hash of the content (request/response)
     * @param metadata Additional metadata (JSON string)
     * @return index of the newly added log entry
     */
    function recordLog(
        string memory userIdentifier,
        string memory actionType,
        string memory contentHash,
        string memory metadata
    ) public returns (uint256) {
        // Create a new log entry
        LogEntry memory newLog = LogEntry({
            recorder: msg.sender,
            userIdentifier: userIdentifier,
            actionType: actionType,
            contentHash: contentHash,
            metadata: metadata,
            timestamp: block.timestamp
        });

        // Add to logs array
        logs.push(newLog);
        uint256 logIndex = logs.length - 1;

        // Update mappings
        userLogs[userIdentifier].push(logIndex);
        hashToLogIndex[contentHash] = logIndex;

        // Emit event
        emit LogRecorded(userIdentifier, actionType, contentHash, block.timestamp);

        return logIndex;
    }

    /**
     * @dev Get a specific log entry by index
     * @param index The index of the log entry
     * @return The log entry details
     */
    function getLog(uint256 index) public view returns (
        address recorder,
        string memory userIdentifier,
        string memory actionType,
        string memory contentHash,
        string memory metadata,
        uint256 timestamp
    ) {
        require(index < logs.length, "Index out of bounds");
        
        LogEntry memory log = logs[index];
        
        return (
            log.recorder,
            log.userIdentifier,
            log.actionType,
            log.contentHash,
            log.metadata,
            log.timestamp
        );
    }

    /**
     * @dev Get all log indices for a specific user
     * @param userIdentifier The user identifier
     * @return Array of log indices
     */
    function getUserLogs(string memory userIdentifier) public view returns (uint256[] memory) {
        return userLogs[userIdentifier];
    }

    /**
     * @dev Get the total number of logs
     * @return The total count of logs
     */
    function getLogCount() public view returns (uint256) {
        return logs.length;
    }

    /**
     * @dev Verify if a content hash exists in the audit trail
     * @param contentHash The hash to verify
     * @return exists Whether the hash exists
     * @return index The index of the log if it exists
     */
    function verifyContentHash(string memory contentHash) public view returns (bool exists, uint256 index) {
        index = hashToLogIndex[contentHash];
        
        // If the hash doesn't exist, the index will be 0 (default value)
        // We need to check if the index is valid and the content hash matches
        exists = index < logs.length && 
                keccak256(bytes(logs[index].contentHash)) == keccak256(bytes(contentHash));
        
        return (exists, index);
    }
} 