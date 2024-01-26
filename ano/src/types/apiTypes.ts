// types/apiTypes.ts

export interface IRLEvent {
    ID: number
    title: string
    description: string
    img: string
    createdAt: number
    updatedAt: number
    // link to event
}

export interface News {
    ID: number
    Title: string
    Description: string
    ImagePath: string
    CreatedAt: number
    UpdatedAt: number
    // link
}

export interface Signup {
    ID: number
    FirstName: string
    LastName: string
    PhoneNumber: number
    Email: string
    CreatedAt: number
}