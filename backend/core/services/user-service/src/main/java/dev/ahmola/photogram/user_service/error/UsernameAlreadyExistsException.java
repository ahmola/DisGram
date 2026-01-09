package dev.ahmola.photogram.user_service.error;

import dev.ahmola.photogram.common_domain.error.ApiException;
import dev.ahmola.photogram.common_domain.error.ErrorCode;
import lombok.extern.slf4j.Slf4j;

@Slf4j
public class UsernameAlreadyExistsException extends ApiException {
    public UsernameAlreadyExistsException() {
        super(ErrorCode.DUPLICATE_ENTITY);
        log.error("Duplicated User");
    }
}
