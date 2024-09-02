import { useEffect, useState } from "react";
import { Note } from "../../types/NoteTypes";
import { fetchNotes } from "../../requests/NoteRequests";
import NoteCard from "./NoteCard";

export default function DesktopSidebar() {
    const [notes, setNotes] = useState<Note[]>([]);

    useEffect(() => {
        fetchNotes()
            .then((data) => setNotes(data));
    })

    return (
        <div className="flex flex-col h-full">
            {/* Sidebar content is in here */}
            <div className="flex-grow overflow-y-auto">
                {/* First, just a section with a M on the left, "Notes" in the center, and "N" on the right */}
                <div className="bg-orange-700 w-full h-12 flex flex-row justify-between items-center">
                    <div className="text-white text-xl font-bold">M</div>
                    <div className="text-white text-xl font-bold">Notes</div>
                    <div className="text-white text-xl font-bold">N</div>
                </div>
                <div className="bg-blue-700 w-full h-10 flex flex-col justify-center items-center">
                    <input type="text" className="bg-green-700 w-full h-full text-center" placeholder="Testing"></input>
                </div>
    
                {/* Then, a list of notes */}
                <div className="overflow-y-auto">
                    {notes.map((note) => (
                        <NoteCard key={note.id} note={note} />
                    ))}
                </div>
            </div>
    
            {/* At the very bottom of the screen, a small display showing current connection status and a GitHub link */}
            <div className="bg-yellow-500 w-full h-4 flex flex-row justify-between items-center">
                {/* Show a green badge showing connected on the left, and a GitHub link on the right */}
                <div className="flex flex-row items-center pl-2">
                    <div className="bg-red-500 w-2 h-2 rounded-full"></div>
                    <span className="text-xs pl-1">Connecting...</span>
                </div>
                <a href="" className="text-white text-xs pr-2">GitHub</a>
            </div>
        </div>
    )
}