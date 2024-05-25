import { User } from "@/types/go/database";
import { api } from "./helpers";

export interface createUploadImageBody {
    title?: string;
    toUserId: string;
    mimeType: string;
}

export interface createUploadResponse {
    title: string;
    guid: string;
    userId: number;
    toUserId: number;
    isUploaded: boolean;
    expiryAt: string;
    mimeType: string;
}

export const uploadImage = async (file: Blob, toUser: User) => {
    const imageRequestBody: createUploadImageBody = {
        toUserId: toUser.id,
        mimeType: "image/jpeg",
    };
    const imageRequest = await api
        .post<createUploadResponse>("/create-upload", imageRequestBody)
        .catch((err) => {
            // logger
            return null;
        });

    if (imageRequest === null) return Promise.reject();

    // let tst =

    // Create actual image
    const buff = await file.arrayBuffer();
    const result = await api
        .post<boolean>(`/upload/${imageRequest.data.guid}`, buff, {
            headers: {
                "Content-Type": file.type,
            },
        })
        .catch((err) => null);

    if (result === null) return Promise.reject();

    return {
        result: result,
        guid: imageRequest,
    };
};
