import { Note } from "../../types/NoteTypes"

export default function NoteCard({note}: {note: Note}) {
    return (
        <div className="bg-slate-400 border-b-red-600 border-b-2 w-full flex flex-col justify-center items-start p-2">
            {/* Title on left side, possible pin icon */}
            <div className="flex flex-row justify-between w-full">
                <div className="text-white text-lg font-bold">{note.title}</div>
                {note.pinned ? <button className="text-white text-lg font-bold">P</button> : <></>}
            </div>
            <div className="text-white text-sm">{note.subtitle}</div>
            {/*Scrollable div for labels */}
            <div className="flex flex-row overflow-x-auto no-scrollbar space-x-1 mt-2">
                {note.labels.map((label) => (
                    <div key={label.id} className={`flex-shrink-0 bg-red-500 text-white text-xs rounded-md`}>
                        <span className="p-1">{label.name}</span>
                    </div>
                ))}
            </div>
        </div>
    )
}