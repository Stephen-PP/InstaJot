# Updated Encrypted Notes App: Models and API Design

## Updated Models

### User
- id: UUID (primary key)
- username: String (unique)
- email: String (encrypted, unique)
- password_hash: String
- salt: String (used for key derivation)
- created_at: Timestamp
- updated_at: Timestamp

### Note
- id: UUID (primary key)
- user_id: UUID (foreign key to User)
- encrypted_content: Text
  - Structure when decrypted:
    ```json
    {
      "title": String,
      "content": String,
      "created_at": Timestamp,
      "updated_at": Timestamp,
      "tags": Array<String>
    }
    ```
- version: Integer
- is_deleted: Boolean
- server_created_at: Timestamp
- server_updated_at: Timestamp

### UserDevice
- id: UUID (primary key)
- user_id: UUID (foreign key to User)
- device_name: String
- last_synced_at: Timestamp

### SyncLog
- id: UUID (primary key)
- user_id: UUID (foreign key to User)
- note_id: UUID (foreign key to Note)
- device_id: UUID (foreign key to UserDevice)
- action: String (e.g., "create", "update", "delete")
- timestamp: Timestamp

### RecoveryKey (Optional)
- id: UUID (primary key)
- user_id: UUID (foreign key to User)
- encrypted_recovery_key: Text
- created_at: Timestamp