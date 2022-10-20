import backend from "@/helpers/backend"
import type { BookStats, Book } from '@/types';


export const getBookstats = () => {

    return backend<BookStats>({
        url: '/api/book/stats',
        method: 'get'
    })
}

interface IndexBook {
    random: Book[],
    recent: Book[],
}


export const bookIndex = () => {
    //<BookStats>
    return backend<IndexBook>({
        url: '/api/book/index',
        method: 'get'
    })
}

export const getBook = (id: string | string[]) => {

    return backend<Book>({
        url: `/api/books/${id}`,
        method: 'get'
    })
}

// index(): Promise < any > {
//     return backend.get("/api/book/index")
// }

// get(id: string | string[]): Promise < any > {
//     return backend.get(`/api/books/${id}`)
// }

// update(id: string | number, data: any): Promise < any > {
//     return backend.put(`/api/books/${id}`, data)
// }

// download(id: string): Promise < any > {
//     return backend.get(`/api/books/${id}/download`)
// }

// search(query: any): Promise < any > {
//     return backend.get("/api/book/search")
// }

// random(max: number): Promise < any > {
//     return backend.get("/api/book/random")
// }

// recent(max: number): Promise < any > {
//     return backend.get("/api/book/recent")
// }