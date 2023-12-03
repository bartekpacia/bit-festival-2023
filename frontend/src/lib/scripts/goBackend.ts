export enum Placements {
    A1 = "A1",
    A2 = "A2",
    B2 = "B2",
    C = "C",
    E = "E",
}

export interface GoReq {
    ampacity: number
    maxPower: number
    veinsUnderLoad: number
    placements: number
    temperature: number
}

export interface GoRsp {
    cableType: string,
    veinsNumber: number,
    crossSection: number,
    insulationType: string,
    placements: Placements,
    temperature: number
}