package com.example.demo.security

import com.example.demo.repository.UserRepository
import org.springframework.security.core.authority.SimpleGrantedAuthority
import org.springframework.security.core.userdetails.UserDetails
import org.springframework.security.core.userdetails.UserDetailsService
import org.springframework.security.core.userdetails.UsernameNotFoundException
import org.springframework.stereotype.Service

@Service
class CustomUserDetailsService(
    private val userRepository: UserRepository
) : UserDetailsService {

    override fun loadUserByUsername(username: String): UserDetails {
        println("🔍 DB 사용자 조회: $username")

        val user = userRepository.findByUsername(username)
            .orElseThrow {
                println("❌ 사용자 없음: $username")
                UsernameNotFoundException("User not found: $username")
            }

        println("✅ 사용자 찾음: ${user.username}")

        return org.springframework.security.core.userdetails.User(
            user.username,
            user.password,
            listOf(SimpleGrantedAuthority("ROLE_${user.role}"))
        )
    }
}
