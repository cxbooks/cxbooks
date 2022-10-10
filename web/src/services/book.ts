import backend from "@/helpers/backend"


class BookService {

    index(): Promise<any> {
        return backend.get("/api/book/index")
    }

    get(id: string|string[]): Promise<any> {
        return backend.get(`/api/books/${id}`)
    }

    update(id: string | number, data: any): Promise<any> {
        return backend.put(`/api/books/${id}`,data)
    }

    download(id: string): Promise<any> {
        return backend.get(`/api/books/${id}/download`)
    }

    search(query: any): Promise<any> {
        return backend.get("/api/book/search")
    }

    random(max: number): Promise<any> {
        return backend.get("/api/book/random")
    }

    recent(max: number): Promise<any> {
        return backend.get("/api/book/recent")
    }

}

export default new BookService();