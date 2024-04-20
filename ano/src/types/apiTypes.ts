// types/apiTypes.ts

export interface IRLEvent {
    ID: number
    Title: string
    Description: string
    ImagePath: string
    Date: string
    CreatedAt: number
    UpdatedAt: number
}

export interface News {
    ID: number
    Title: string
    Description: string
    ImagePath: string
    CreatedAt: number
    UpdatedAt: number
}

export interface Signup {
    ID: number
    FirstName: string
    LastName: string
    PhoneNumber: string
    Email: string
    CreatedAt: number
}

export interface Photo {
    ID: number
    ImagePath: string
}

export interface Athlete {
    ID: number
    Title: string
    Description: string
    ImagePath: string
    CreatedAt: number
}