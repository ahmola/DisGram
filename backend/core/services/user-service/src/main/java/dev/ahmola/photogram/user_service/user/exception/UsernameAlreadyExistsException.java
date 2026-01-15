package dev.ahmola.photogram.user_service.user.exception;

import dev.ahmola.photogram.common_domain.error.ApiException;
import dev.ahmola.photogram.common_domain.error.ErrorCode;
import lombok.extern.slf4j.Slf4j;

@Slf4j
public class UsernameAlreadyExistsException extends ApiException {
    public UsernameAlreadyExistsException(String username) {
        super(ErrorCode.DUPLICATE_ENTITY);
        log.error("Duplicated Username : {}", username);
    }
}
