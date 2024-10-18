## Updated API Routes

#### Authentication Routes
```
POST /api/auth/register
 Registers a new user. Expects username, email, and password. Returns user ID and authentication token.

POST /api/auth/login
 Authenticates a user. Expects username/email and password. Returns user ID and authentication token.

POST /api/auth/logout
 Invalidates the current user's authentication token.
```

#### User Routes
```
GET /api/users/me
 Retrieves the current user's profile information.

PATCH /api/users/me
 Updates the current user's profile information. Expects updated fields in the request body.
```

#### Notes Routes
```
POST /api/notes
 Creates a new note. Expects encrypted content string in the request body.

GET /api/notes
 Retrieves all notes for the current user. Returns a list of note IDs and server timestamps.

GET /api/notes/{note_id}
 Retrieves a specific note by ID. Returns the encrypted content string.

PUT /api/notes/{note_id}
 Updates a note's content. Expects the full encrypted content string in the request body.

DELETE /api/notes/{note_id}
 Marks a note as deleted (soft delete unless in trash already).
```

#### Sync Routes
```
GET /api/notes/{note_id}/sync
 Retrieves the sync status for a specific note. Returns the note's current version on the server and last_synced_at timestamp.

POST /api/notes/{note_id}/sync
 Synchronizes a specific note. Expects the encrypted note content, client version, and last synced timestamp in the request body. Returns the updated note data if successful, or conflict information if a conflict is detected.
```

#### Devices Routes
```
POST /api/devices
 Registers a new device for the current user. Expects device name in the request body.

GET /api/devices
 Retrieves a list of all devices registered to the current user.

DELETE /api/devices/{device_id}
 Removes a device from the user's account.
```

#### Recovery Routes (Optional)
```
POST /api/recovery/key
 Sets or updates the user's encrypted recovery key. Expects the encrypted recovery key in the request body.

POST /api/recovery/account
 Initiates the account recovery process. Expects recovery key and new password in the request body.
```