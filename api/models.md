# Encrypted Notes App: Complete Schema

## User
- id: UUID (primary key)
- username: String (unique)
- email: String (encrypted, unique)
- password_hash: String
- salt: String (used for key derivation)
- created_at: Timestamp
- updated_at: Timestamp

## Note
- id: UUID (primary key)
- user_id: UUID (foreign key to User)
- encrypted_metadata: Text
  - Structure when decrypted:
    ```json
    {
      "title": String,
      "created_at": Timestamp,
      "updated_at": Timestamp,
      "tags": Array<String>
    }
    ```
- version: Integer
- is_deleted: Boolean
- server_created_at: Timestamp
- server_updated_at: Timestamp

## NoteBlock
- id: UUID (primary key)
- note_id: UUID (foreign key to Note)
- encrypted_content: Text
  - Structure when decrypted:
    ```json
    {
      "type": String (e.g., "paragraph", "heading", "list", "code", "image"),
      "content": String,
      "metadata": {
        "alignment": String,
        "styles": Array<String>,
        "level": Integer (for headings),
        "language": String (for code blocks),
        "alt_text": String (for images)
      }
    }
    ```
- order: Integer
- version: Integer
- created_at: Timestamp
- updated_at: Timestamp

## UserDevice
- id: UUID (primary key)
- user_id: UUID (foreign key to User)
- device_name: String
- last_synced_at: Timestamp

## SyncLog
- id: UUID (primary key)
- user_id: UUID (foreign key to User)
- note_id: UUID (foreign key to Note)
- device_id: UUID (foreign key to UserDevice)
- action: String (e.g., "create", "update", "delete")
- timestamp: Timestamp

## RecoveryKey (Optional)
- id: UUID (primary key)
- user_id: UUID (foreign key to User)
- encrypted_recovery_key: Text
- created_at: Timestamp
```

Explanation of the schema and how encryption affects it:

1. User Model:
   - The `salt` is used in the key derivation process, unique per user.
   - `email` is encrypted to enhance privacy.
   - `password_hash` is stored for authentication, but the actual encryption key is never stored.

2. Note Model:
   - `encrypted_metadata` contains the note's metadata (title, timestamps, tags) encrypted as a JSON string.
   - `version` helps with conflict resolution during syncing.
   - `server_created_at` and `server_updated_at` are used for sorting and syncing but don't reveal actual note modification times.

3. NoteBlock Model:
   - Each block's content and metadata are encrypted separately, allowing for efficient updates to individual blocks.
   - `order` maintains the sequence of blocks within a note.
   - `version` helps with block-level conflict resolution.

4. UserDevice Model:
   - Keeps track of different devices a user uses to access their notes.
   - Helps in managing sync states across devices.

5. SyncLog Model:
   - Maintains a record of sync actions for conflict resolution and auditing.
   - Helps in implementing efficient incremental sync.

6. RecoveryKey Model (Optional):
   - Stores an encrypted recovery key that can be used to recover notes if a user forgets their password.
   - The recovery key is encrypted with a separate passphrase or through other secure means.

Encryption Implications:
1. All sensitive data (note content, metadata, email) is encrypted client-side before being sent to the server.
2. The encryption key is derived from the user's password using the stored salt, but the key itself is never stored.
3. Each note block is encrypted separately, allowing for efficient updates and sync of individual blocks.
4. The server only sees encrypted data and cannot decrypt it, ensuring zero-knowledge principle.
5. Timestamps in the Note and NoteBlock models are for server-side operations. Actual creation and modification times are stored in the encrypted metadata.

This schema design allows for:
- Secure, end-to-end encrypted storage of notes
- Efficient syncing across multiple devices
- Granular updates to note content via the block structure
- Maintaining the zero-knowledge principle while still allowing for features like search and sorting (through encrypted metadata)

Considerations:
- Implement proper key derivation (e.g., PBKDF2, Argon2) on the client-side.
- Ensure all encryption/decryption happens client-side.
- Implement a secure method for key rotation when a user changes their password.
- Consider implementing a secure backup or recovery system, as forgotten passwords will result in permanent data loss in a true zero-knowledge system.

Would you like me to elaborate on any specific part of this schema or discuss how certain operations would work with this encrypted structure?