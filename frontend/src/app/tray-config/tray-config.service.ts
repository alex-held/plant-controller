import {Injectable} from "@angular/core";
import {Observable} from "rxjs";
import {TrayConfig} from "./tray-config.model";
import {HttpClient} from "@angular/common/http";
import {environment} from "../../environments/environment";

@Injectable({providedIn: "any"})
export class TrayConfigService {
    private baseUrl = `${environment.apiUrl}/v1/trayconfig/`;

    constructor(private httpClient: HttpClient) {
    }

    create(growConfig: TrayConfig): Observable<TrayConfig> {
        const url = this.baseUrl
        console.log('POST | ' + url);
        return this.httpClient.post<TrayConfig>(url, growConfig)
    }

    read(): Observable<TrayConfig[]> {
        const url = this.baseUrl
        console.log('GET | ' + url);
        return this.httpClient.get<TrayConfig[]>(url)
    }

    getById(id: number): Observable<TrayConfig> {
        const url = `${this.baseUrl}${id}`
        console.log('GET | ' + url);
        return this.httpClient.get<TrayConfig>(url);
    }


    update(growConfig: TrayConfig): Observable<TrayConfig> {
        const url = `${this.baseUrl}${growConfig["Id"]}`
        console.log('PUT | ' + url);
        return this.httpClient.put<TrayConfig>(url, growConfig);
    }

    delete(id: number): Observable<any> {
        const url = `${this.baseUrl}${id}`
        console.log('DELETE | ' + url);
        return this.httpClient.delete<any>(url);
    }

}
