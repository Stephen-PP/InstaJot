export interface Note {
    id: number
    title: string
    subtitle: string
    content: string
    pinned: boolean
    labels: NoteLabel[]
    createdAt: number
    updatedAt: number
}

export interface NoteLabel {
    id: number
    name: string
    color: string
}