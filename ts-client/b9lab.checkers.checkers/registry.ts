import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateGame } from "./types/checkers/tx";
import { MsgPlayMove } from "./types/checkers/tx";
import { MsgRejectGame } from "./types/checkers/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/b9lab.checkers.checkers.MsgCreateGame", MsgCreateGame],
    ["/b9lab.checkers.checkers.MsgPlayMove", MsgPlayMove],
    ["/b9lab.checkers.checkers.MsgRejectGame", MsgRejectGame],
    
];

export { msgTypes }