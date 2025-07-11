package com.example.demo.controller

import com.example.demo.domain.User
import com.example.demo.dto.AuthRequest
import com.example.demo.dto.AuthResponse
import com.example.demo.repository.UserRepository
import com.example.demo.security.JwtProvider
import org.springframework.http.ResponseEntity
import org.springframework.security.authentication.AuthenticationManager
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken
import org.springframework.security.crypto.password.PasswordEncoder
import org.springframework.web.bind.annotation.*

@RestController
@RequestMapping("/api/auth")
class AuthController(
    private val userRepository: UserRepository,
    private val passwordEncoder: PasswordEncoder,
    private val jwtProvider: JwtProvider,
    private val authenticationManager: AuthenticationManager
) {

    @PostMapping("/signup")
    fun signup(@RequestBody request: AuthRequest): ResponseEntity<String> {
        if (userRepository.findByUsername(request.username).isPresent) {
            return ResponseEntity.badRequest().body("User already exists")
        }

        val user = User(
            username = request.username,
            password = passwordEncoder.encode(request.password),
            role = "USER"
        )
        userRepository.save(user)
        return ResponseEntity.ok("User registered successfully")
    }

    @PostMapping("/login")
    fun login(@RequestBody request: AuthRequest): ResponseEntity<AuthResponse> {
        println("🔐 로그인 요청: ${request.username}")

        try {
            val authToken = UsernamePasswordAuthenticationToken(request.username, request.password)
            val authResult = authenticationManager.authenticate(authToken)

            println("✅ 인증 성공: ${authResult.isAuthenticated}")

            val user = userRepository.findByUsername(request.username).orElseThrow {
                RuntimeException("사용자 조회 실패")
            }

            val token = jwtProvider.generateToken(user.username, user.role)
            return ResponseEntity.ok(AuthResponse(token))

        } catch (e: Exception) {
            println("❌ 인증 실패: ${e.message}")
            return ResponseEntity.status(401).build()
        }
    }

}
