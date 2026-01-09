package dev.ahmola.photogram.common_domain.error;

import lombok.Getter;
import lombok.extern.slf4j.Slf4j;

@Slf4j
@Getter
public class ApiException extends RuntimeException{
    private final ErrorCode errorCode;
    public ApiException(ErrorCode errorCode){
        super(errorCode.getMessage());
        this.errorCode = errorCode;
        log.error("API Exception Occured");
    }
}
