import {Injectable} from "@angular/core";
import {Observable} from "rxjs";
import {HttpClient} from "@angular/common/http";
import {environment} from "../../environments/environment";
import {Tray} from "./tray.model";

@Injectable({providedIn: "any"})
export class TrayService {
    private baseUrl = `${environment.apiUrl}/v1/tray/`;

    constructor(private httpClient: HttpClient) {
    }

    create(growConfig: Tray): Observable<Tray> {
        const url = this.baseUrl
        console.log('POST | ' + url);
        return this.httpClient.post<Tray>(url, growConfig)
    }

    read(): Observable<Tray[]> {
        const url = this.baseUrl
        console.log('GET | ' + url);
        return this.httpClient.get<Tray[]>(url)
    }

    getById(id: number): Observable<Tray> {
        const url = `${this.baseUrl}${id}`
        console.log('GET | ' + url);
        return this.httpClient.get<Tray>(url);
    }


    update(growConfig: Tray): Observable<Tray> {
        const url = `${this.baseUrl}${growConfig["Id"]}`
        console.log('PUT | ' + url);
        return this.httpClient.put<Tray>(url, growConfig);
    }

    delete(id: number): Observable<any> {
        const url = `${this.baseUrl}${id}`
        console.log('DELETE | ' + url);
        return this.httpClient.delete<any>(url);
    }

}
