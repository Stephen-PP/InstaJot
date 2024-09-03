import { useEffect, useState } from "react";
import { Note } from "../../types/NoteTypes";
import { fetchNotes } from "../../requests/NoteRequests";
import NoteCard from "./NoteCard";

export default function DesktopSidebar() {
    const [notes, setNotes] = useState<Note[]>([]);

    useEffect(() => {
        fetchNotes()
            .then((data) => setNotes(data));
    }, []);

    return (
        <div className="flex flex-col h-full hidden lg:block relative">
            {/* Sidebar content is in here */}
            <div className="flex-grow">
                {/* First, just a section with "Notes" in the center, and "N" on the right */}
                <div className="bg-orange-700 w-full h-12 flex flex-row justify-between items-center relative">
                    <div></div>
                    <div className="text-white text-xl font-bold absolute left-1/2 transform -translate-x-1/2">Notes</div>
                    <div className="text-white text-xl font-bold pr-2">N</div>
                </div>
                <div className="bg-blue-700 w-full h-10 flex flex-col justify-center items-center">
                    <input type="text" className="bg-green-700 w-full h-full text-center" placeholder="Testing"></input>
                </div>
    
                {/* Then, a list of notes */}
                <div>
                    {notes.map((note) => (
                        <NoteCard key={note.id} note={note} />
                    ))}
                </div>
            </div>
    
            {/* At the very bottom of the sidebar, a small display showing current connection status and a GitHub link */}
            <div className="bg-yellow-500 w-full h-6 flex flex-row justify-between items-center absolute bottom-0">
                {/* Show a green badge showing connected on the left, and a GitHub link on the right */}
                <div className="flex flex-row items-center pl-2">
                    <div className="bg-red-500 w-2 h-2 rounded-full"></div>
                    <span className="text-xs pl-1">Connecting...</span>
                </div>
                <div className="flex flex-row items-center pr-2">
                    <a href="" className="text-white text-xs">Settings</a>
                    <a href="" className="text-white text-xs pl-2">GitHub</a>
                </div>
            </div>
        </div>
    )
}