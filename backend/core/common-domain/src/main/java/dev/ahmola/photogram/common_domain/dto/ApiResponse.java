package dev.ahmola.photogram.common_domain.dto;

import org.springframework.http.HttpStatus;

import dev.ahmola.photogram.common_domain.error.ErrorCode;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Getter;

@Getter
@AllArgsConstructor
@Builder
public class ApiResponse<T> {
    private final boolean success;
    private final T data;
    private final HttpStatus status;
    private final String message;

    public static <T> ApiResponse<T> success(T data){
        return new ApiResponse<T>(true, data, null, null);
    }

    public static ApiResponse<?> fail(ErrorCode errorCode){
        return new ApiResponse<>(false, null, errorCode.getStatus(), errorCode.getMessage());
    }

}
