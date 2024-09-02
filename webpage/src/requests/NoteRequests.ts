import { Note } from "../types/NoteTypes";

export async function fetchNotes() {
    const notes: Note[] = [{
        id: 1,
        title: 'Test Note',
        subtitle: 'This note has some test..',
        content: 'This note has some test content I guess',
        pinned: false,
        labels: [],
        createdAt: 0,
        updatedAt: 0
    },
    {
        id: 2,
        title: 'Test Note 2',
        subtitle: 'This note has some test..',
        content: 'This note has some test content I guess',
        pinned: false,
        labels: [{
            id: 1,
            name: 'Test Label',
            color: 'blue'
        },
        {
            id: 2,
            name: 'Test Label 2',
            color: 'red'
        },
        {
            id: 3,
            name: 'Test Label 3',
            color: 'red'
        },
        {
            id: 4,
            name: 'Test Label 4',
            color: 'red'
        },
        {
            id: 5,
            name: 'Test Label 5',
            color: 'red'
        }],
        createdAt: 0,
        updatedAt: 0
    },
    {
        id: 3,
        title: 'Test Note 3',
        subtitle: 'This note has some test..',
        content: 'This note has some test content I guess',
        pinned: true,
        labels: [],
        createdAt: 0,
        updatedAt: 0
    }]

    return notes;
}